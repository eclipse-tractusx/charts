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
	repos, _, _ := getOrgRepos(ctx, gitOwner, client)

	err := initHelmRepoIndex(indexFile)
	if err != nil {
		log.Fatal(err)
	}

	for i, gitRepo := range repos {
		var gitRepoHelmIndex string

		//check gh pages configured for repo
		_, response, _ := client.Repositories.GetBranch(ctx, gitOwner, *gitRepo.Name, "gh-pages", false)

		if response.StatusCode == 200 {
			fmt.Println(i, *gitRepo.Name)
			gitRepoHelmIndex = downloadProductHelmRepoIndex(gitOwner, *gitRepo.Name)
			buildHelmRepoIndex(indexFile, gitRepoHelmIndex)
		}
	}
}

func sayHelloTo(input string) {
	fmt.Printf("Hello %v", input)
}

func getAuthenticatedClient(ctx context.Context, gitToken string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gitToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

func getOrgRepos(ctx context.Context, gitOwner string, client *github.Client) ([]*github.Repository, *github.Response, error) {
	opt := &github.RepositoryListByOrgOptions{
		Type: "all",
	}
	repos, response, err := client.Repositories.ListByOrg(ctx, gitOwner, opt)

	return repos, response, err
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

func downloadProductHelmRepoIndex(gitOwner string, gitRepo string) string {
	fileName := gitRepo + "-index.yaml"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	response, _ := http.Get(fmt.Sprintf("https://raw.githubusercontent.com/%v/%v/gh-pages/index.yaml", gitOwner, gitRepo))

	if response.StatusCode == 200 {
		_, _ = io.Copy(file, response.Body)

		// Close response body
		err := response.Body.Close()
		if err != nil {
			log.Fatal(err)
		}

		// Close file
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	err = file.Close()
	return fileName
}

func buildHelmRepoIndex(indexFile string, mergeIndexFile string) {
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
