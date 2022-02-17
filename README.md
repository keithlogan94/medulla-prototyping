

# Prototype

### How to Work with this Prototype

This prototype implements a web server that listens at /

When any GET request is sent to the server /some-file.html. 
The service will check for that file /some-file on an s3 bucket, if it exists 
the requesting user will be redirected to the s3 bucket using a `Presigned URL`.

Any path can be specified

`/some-file.txt`

`/bundle.js`

Multiple directories can be specified

`/some-directory/bundle.js`

As long as the path exists in the s3 bucket

the user will be redirected to that file with a presigned url.


### How to Setup this Prototype

First configure your s3 bucket or minio instance.

Once your s3 bucket is configured, then capture the access key id, and secret access key.

Create a directory in the s3 bucket which you want to serve files out of.

Add a file into that directory which you want to test.

You should have all the info needed to fill out the `tye.yaml` file with your settings.

Update the following env vars in the `tye.yaml` file

`MINIO_ENDPOINT` 

`MINIO_ACCESS_KEY`

`MINIO_SECRET_ACCESS_KEY`

```
- name: S3_BUCKET_FOLDER
  value: "[folder]"
    - name: MINIO_ENDPOINT
      value: "[endpoint]"
    - name: MINIO_ACCESS_KEY
      value: "[access-key]"
    - name: MINIO_SECRET_ACCESS_KEY
      value: "[secret-access-key]"
```
