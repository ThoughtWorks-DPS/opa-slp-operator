The opa-sl-operator deploys the Styra Local Plane configured to match the Styra system pattern where all opa sidecars in a namespace share the same styra System and it's associated policy configuration.  


**example:**  

Assume a development team at book-info company is named the `publications` team since they own the product domain of _publications_ within the book-info SaaS. The publications team's path to production is through the following environments:  

dev => qa => production  

and that these environments equate to the matching namespaces of:  

publications-dev => publications-qa => publications-prod  

The publications team has a single git repository named `publications-authorization-policies` in which their team specific rego policies are maintained. There may be a styra Stack or shared libraries containing additional policies that will be included in their policy bundle.

The team would create matching Systems in their organization's Styra tenant, tied to the associated branches in the publications-authorization-policies_ repo:  

Systems:  
- publications-dev => dev branch  
- publications-qa => qa branch  
- publications-prod => prod branch  

Typically, it is recommended that the policy release pipeline for the policies repo follow a pattern of:  

The team develops their policies within the repo using trunk-based-development against the `main` (or `master`) branch. The repo pipeline will then merge all changes pushed to `main` into the dev branch.  

To trigger a release to production, the team tags the main branch with the appropriate semantic version (or other versioning scheme as desired). Tagging causes the pipeline to first merge with the QA branch, and then upon approval further merge with the prod branch.  

Each of these branches is pulled in the respective styra System by the DAS github integration.  

> Within Styra, the publications-dev system can either immediately publish all such changes out to the slp's pulling from publications-dev or wait for a human to 'Publish' in the DAS UI depending on the workflow desired by the team. They may wish to take advantage of Styra's playback features before committing changes to all or only some of the environments.  
