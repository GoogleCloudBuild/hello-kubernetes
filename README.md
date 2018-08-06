# hello-kubernetes
Example end-to-end build, test, and deploy pipelines for Kubernetes, using Google Cloud Build.

# Instructions

If you are new to Google Cloud Build, we recommend you start by visiting the [manage resources page](https://console.cloud.google.com/cloud-resource-manager) in the Cloud Console, [learn how to enable billing](https://cloud.google.com/billing/docs/how-to/modify-project), [enable the Cloud Build API](https://console.cloud.google.com/flows/enableapi?apiid=cloudbuild.googleapis.com), and [install the Cloud SDK](https://cloud.google.com/sdk/docs/). You may also wish to review the [Quickstart for Go](https://cloud.google.com/cloud-build/docs/quickstart-go).

## Continuous integration testing

Google Cloud Build supports continuous integration testing, with a test running automatically every time a Pull Request is created or updated in GitHub.

To try this out for yourself with a simple Go application which prints a greeting:

* Fork this repo on GitHub by clicking the "Fork" icon in the top right-hand corner of the screen.
* Open the [Google Cloud Build GitHub app](https://github.com/apps/google-cloud-build), click Install, and follow the setup instructions.  Select the `hello-kubernetes` repository you just created.
* Back in your forked repo in GitHub, click to open `main_test.go`.  You will see some Go tests which make sure that the greeting matches the name provided.
* Click the pencil icon in the top right-hand corner of the screen.
* Edit one of the greetings so that the greeting does not match the name.
* At the bottom of the screen, choose "Create a new branch and start a pull request".  Click "Propose file change".
* On the page which opens, click "Create pull request".
* You should see a message: "All checks have failed".  If you wish, you can see the build log by clicking the "Details" button and then selecting the link in small print at the bottom of the screen which says "View more details on Google Cloud Build".
* Next, choose the "Files changed" tab, and click the pencil icon again.
* Make another edit so that the greetings now match the names.
* Click the "Commit changes" button at the bottom of the screen.
* Go back to the "Conversation" tab.
* You should see that tests pass, and a green check is now returned.
* Click "Merge pull request" to commit your change to `master`.

## Automated deployment

Google Cloud Build can be used to deploy to Kubernetes automatically whenever you merge to a particular Git branch or tag.  Here, we're going to deploy to Google Kubernetes Engine (GKE) whenever a merge to `master` occurs.

Before we start, we need to enable the GKE API, create a Kubernetes cluster for testing, and grant deployment permission to the Cloud Build service account.  To do this, copy and paste the following lines into your terminal:

```sh
gcloud services enable container.googleapis.com
gcloud container clusters create my-cluster --zone=us-central1-f
export PROJECT=$(gcloud info --format='value(config.project)')
export PROJECT_NUMBER=$(gcloud projects describe $PROJECT --format 'value(projectNumber)')
export CB_SA_EMAIL=$PROJECT_NUMBER@cloudbuild.gserviceaccount.com
gcloud projects add-iam-policy-binding $PROJECT --member=serviceAccount:$CB_SA_EMAIL --role='roles/container.developer'  
```

Then:

* Visit the [Build Triggers page](https://console.cloud.google.com/cloud-build/triggers) in the Cloud Console.
* Click "Create trigger" (if you are new to Cloud Build) or "Add trigger" (if you have existing triggers configured).
* Choose source GitHub and click Continue.  Follow the authorization flow to grant Google access to your code in GitHub.
* Choose the `hello-kubernetes` repo, check the consent box, and click Continue.
* Choose a name for your trigger, e.g. "Deploy to Kubernetes"
* Under Branch (regex), enter `master`.
* Under Build configuration, choose "cloudbuild.yaml", and enter filename `deploy.yaml`.
* Click "Create trigger".
* Now, return to your fork of `hello-kubernetes` in GitHub, and click `main.go` to open the file on the master branch.
* Click the pencil icon on the top right-hand corner to edit.
* On line 16, change the string inside `Greet("...")` to be your name.
* At the bottom of the page, ensure "Commit directly to the master branch" is selected, and click "Commit changes".
* A build, push, and deploy is now running in the background.  This normally takes about 20 seconds.  If you wish you can inspect logs in the [Build History](https://console.cloud.google.com/cloud-build/builds) page in the Cloud Console.
* To see your new code in action, run this command:

```sh
kubectl logs -l run=hello-kubernetes
```

* Kubernetes should now greet you by name.  If you see a message that the container is waiting to start, you may have to wait a little longer.
* To delete your test Kubernetes cluster, run:

```sh
gcloud container clusters delete my-cluster
```
