---
title: 'File Upload for Static Websites using AWS Lambdas'
date: 2020-07-12
permalink: /posts/2020/07/file-upload-static-website-AWS-lambda/
tags:
  - tutorials
---


```js
const AWS = require('aws-sdk');
const s3 = new AWS.S3();

exports.handler = async (event) => {
    let request = JSON.parse(event.body);
    let base64String = request.base64String;
    const bucketName = ''staticfileuploads';
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
