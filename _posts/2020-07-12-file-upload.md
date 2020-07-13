---
title: 'File Upload for Static Websites using AWS Lambdas'
date: 2020-07-12
permalink: /posts/2020/07/file-upload-static-website-AWS-lambda/
tags:
  - tutorials
---

I love everything serverless. This website is a static website built with Jekyll. I needed to have a feature to allow users to be able to upload images or files(when submitting forms) and storing them somewhere without building a backend. So, using AWS, I thought of writing a lambda that's connected to an S3 bucket and some javascript code to handle the form uploading.

By doing this, you wouldn't have to worry about scalability, cost, or losing files since it's all handled by AWS and lambdas are really cheap.

## Bucket

First thing you'll need to create an S3 bucket, I named mine "staticfileuploads".

If you want all the files in the bucket to be publicly accessible then we can turn off "block all public access" from the permissions tab.

![Image of permissions](https://yalabsi.com/images/static-upload/1.png)

## Lambda

### Permissions
After you create the bucket head over to lambdas and create a new lambda.
Go for the basic execution role and attach an S3 policy to access the bucket. You can go with a S3:* permission but I'd recommend only giving the least access needed for the IAM role.

### Trigger
We need to add a URL that we can post to, go tho the lambda and click on add a trigger then select API Gateway. I went with open authentication and open CORS but you can go with a scope of only your website. Once you added it, save the URL for the next steps.

### Code
From your terminal, create a new directory and run the command and follow the default steps:
`npm init`

index.js:
```js
const AWS = require('aws-sdk');
const s3 = new AWS.S3();

exports.handler = async (event) => {
    let request = JSON.parse(event.body);
    let base64String = request.base64String;
    const bucketName = 'staticfileuploads'; // Modify your bucket name
    const fileName = request.fileName;
    try {
      const params = {
            Bucket: bucketName,
            Key: fileName,
            Body: Buffer.from(base64String, 'base64')
        };
       const putResult = await s3.putObject(params).promise();
       return "File Uploaded Successfully";
    } catch (error) {
        console.log(error);
        return;
    }
}
```
 I added a fileName parameter but you can always get the type from the base64 encoding and randomize the key to avoid duplicates.

Make sure you install the package `npm install aws-sdk`, after you're done zip the files(including node-modules) and upload them to the lambda using Actions(right side) in the Function code, then upload zip file.

### Test
After its uploaded we're ready to test it. Here's a post request and the response using Postman. I encoded the file using this [website](https://www.base64encode.org/).

![Postman](https://yalabsi.com/images/static-upload/3.png)


## Frontend

For the frontend, you can have use a function to encode the file then use a fetch to send the request.

Javascript:
```js
function getBase64(file) {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => resolve(reader.result);
    reader.onerror = error => reject(error);
  });
}
```
