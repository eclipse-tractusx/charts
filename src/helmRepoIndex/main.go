/********************************************************************************
 * Copyright (c) 2021,2022 Mercedes-Benz Tech Innovation Gmbh
 * Copyright (c) 2021,2022 Contributors to the Eclipse Foundation
 *
 * See the NOTICE file(s) distributed with this work for additional
 * information regarding copyright ownership.
 *
 * This program and the accompanying materials are made available under the
 * terms of the Apache License, Version 2.0 which is available at
 * https://www.apache.org/licenses/LICENSE-2.0.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 ********************************************************************************/

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
	"helm.sh/helm/v3/pkg/repo"
)

func main() {
	var indexFile string
	var gitOwner string
	var gitToken string

	flag.StringVar(&indexFile, "indexFile", "", "Specify Helm Repository index file")
	flag.StringVar(&gitOwner, "owner", "", "Specify GitHub User or Organization")
	flag.StringVar(&gitToken, "token", "", "Specify GitHub Token")
	flag.Parse()

	ctx := context.Background()
	client := getAuthenticatedClient(ctx, gitToken)
	repos, err := getOrgRepos(ctx, gitOwner, client)
	if err != nil {
		log.Fatal(err)
	}

	err = initHelmRepoIndex(indexFile)
	if err != nil {
		log.Fatal(err)
	}

	for _, gitRepo := range repos {
		var gitRepoHelmIndex string

		gitRepoHelmIndex, err = downloadProductHelmRepoIndex(gitOwner, *gitRepo.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		buildHelmRepoIndex(indexFile, gitRepoHelmIndex, *gitRepo.Name)
	}
}

func getAuthenticatedClient(ctx context.Context, gitToken string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gitToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

func getOrgRepos(ctx context.Context, gitOwner string, client *github.Client) ([]*github.Repository, error) {
	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{},
	}

	var allRepos []*github.Repository

	for {
		repos, response, err := client.Repositories.ListByOrg(ctx, gitOwner, opt)
		if err != nil {
			return nil, err
		}
		allRepos = append(allRepos, repos...)
		if response.NextPage == 0 {
			break
		}
		opt.Page = response.NextPage
	}

	return allRepos, nil
}

func initHelmRepoIndex(fileName string) error {
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		newIndexFile := repo.NewIndexFile()
		err = newIndexFile.WriteFile(fileName, 0644)

		if err != nil {
			return err
		}
	}
	return nil
}

func downloadProductHelmRepoIndex(gitOwner string, gitRepo string) (string, error) {
	response, _ := http.Get(fmt.Sprintf("https://raw.githubusercontent.com/%v/%v/gh-pages/index.yaml", gitOwner, gitRepo))
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)

	if response.StatusCode == 200 {
		log.Printf("✅  %v - index.yaml download completed", gitRepo)
		fileName := gitRepo + "-index.yaml"
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}

		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(file)

		_, err = io.Copy(file, response.Body)
		if err != nil {
			log.Fatal(err)
		}
		return fileName, nil
	} else {
		dt := time.Now()
		return "", fmt.Errorf("%v ❌  %v - no index.yaml found", dt.Format("2006/01/02 15:04:05"), gitRepo)
	}
}

func buildHelmRepoIndex(indexFile string, mergeIndexFile string, gitRepo string) {
	repoFile, err := repo.LoadIndexFile(indexFile)
	if err != nil {
		log.Fatal(err)
	}

	// merge index only if file has content and is not from local repository
	if fileStat, _ := os.Stat(mergeIndexFile); fileStat.Size() > 0 && mergeIndexFile != "tractusx-release" {
		newIndex, err := repo.LoadIndexFile(mergeIndexFile)

		if err != nil {
			log.Fatal(err)
		}
		repoFile.Merge(newIndex)
		log.Printf("✅  %v - index.yaml merged into Helm repository DEV", gitRepo)
		repoFile.Generated = time.Now()
		err = repoFile.WriteFile(indexFile, 0644)

		if err != nil {
			log.Fatal(err)
		}
	}
	err = os.Remove(mergeIndexFile)
	if err != nil {
		log.Fatal(err)
	}
}
