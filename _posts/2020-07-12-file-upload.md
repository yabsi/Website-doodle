---
title: 'File Upload for Static Websites using AWS Lambdas'
date: 2020-07-12
permalink: /posts/2020/07/file-upload-static-website-AWS-lambda/
tags:
  - tutorials
---

Working on this static website(It's a Jekyll website with no backend), I realized how nice of an option it is, and was thinking of some of the drawbacks like being able to upload images or files(when submitting forms) and storing them somewhere. Without building a backend, I thought of writing a lambda that's connected to an S3 bucket and some javascript code to handle the form uploading.

First thing you'll need to create an S3 bucket, I named mine "staticfileuploads".

If you want all the files in the bucket to be publicly accessible then we can turn off "block all public access" from the permissions tab.


```js
const AWS = require('aws-sdk');
const s3 = new AWS.S3();

exports.handler = async (event) => {
    let request = JSON.parse(event.body);
    let base64String = request.base64String;
    const bucketName = 'staticfileuploads';
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
