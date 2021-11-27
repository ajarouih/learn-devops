# Cloudfront

[Cheat Sheet - Cloudfront](https://tutorialsdojo.com/amazon-cloudfront)

[geoproximity-routing-vs-geolocation-routing](https://tutorialsdojo.com/latency-routing-vs-geoproximity-routing-vs-geolocation-routing)

[What is Amazon CloudFront?](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Introduction.html)

Amazon CloudFront is a web service that speeds up distribution of your static and dynamic web content, such as .html, .css, .js, and image files, to your users. CloudFront delivers your content through a worldwide network of data centers called edge locations. When a user requests content that you're serving with CloudFront, the request is routed to the edge location that provides the lowest latency (time delay), so that content is delivered with the best possible performance.

- If the content is already in the edge location with the lowest latency, CloudFront delivers it immediately.

- If the content is not in that edge location, CloudFront retrieves it from an origin that you've defined—such as an Amazon S3 bucket, a MediaPackage channel, or an HTTP server (for example, a web server) that you have identified as the source for the definitive version of your content.


## Restricting the geographic distribution of your content

[Restricting the geographic distribution of your content](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/georestrictions.html)

You can use geo restriction, also known as geo blocking, to prevent users in specific geographic locations from accessing content that you're distributing through a CloudFront distribution. To use geo restriction, you have two options:

- Use the CloudFront geo restriction feature. Use this option to restrict access to all of the files that are associated with a distribution and to restrict access at the country level.

- Use a third-party geolocation service. Use this option to restrict access to a subset of the files that are associated with a distribution or to restrict access at a finer granularity than the country level.

## Optimizing Caching and availability

### Optimizing high availability with CloudFront origin failover

[Optimizing high availability with CloudFront origin failover](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/high_availability_origin_failover.html)


- You can set up CloudFront with origin failover for scenarios that require high availability. 
- To get started, you create an origin group with two origins: a primary and a secondary. 
- If the primary origin is unavailable, or returns specific HTTP response status codes that indicate a failure, CloudFront automatically switches to the secondary origin.


## Configuring secure access and restricting access to content

### Using field-level encryption to help protect sensitive data

[Using field-level encryption to help protect sensitive data](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/field-level-encryption.html)

- With Amazon CloudFront, you can enforce secure end-to-end connections to origin servers by using HTTPS
- Field-level encryption adds an additional layer of security that lets you protect specific data throughout system processing so that only certain applications can see it.
- Field-level encryption allows you to enable your users to securely upload sensitive information to your web servers.
- The sensitive information provided by your users is encrypted at the edge, close to the user, and remains encrypted throughout your entire application stack. 
- This encryption ensures that only applications that need the data—and have the credentials to decrypt it—are able to do so.
- To use field-level encryption, when you configure your CloudFront distribution, specify the set of fields in POST requests that you want to be encrypted, and the public key to use to encrypt them. 
- You can encrypt up to 10 data fields in a reques


## Optimizing caching and availability

### Increasing the proportion of requests that are served directly from the CloudFront caches (cache hit ratio)

[Increasing the proportion of requests that are served directly from the CloudFront caches (cache hit ratio)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cache-hit-ratio.html)

- You can improve performance by increasing the proportion of your viewer requests that are served directly from the CloudFront cache instead of going to your origin servers for content. 
- This is known as improving the cache hit ratio.

- Specifying how long CloudFront caches your objects
  - To increase your cache hit ratio, you can configure your origin to add a Cache-Control max-age directive to your objects, and specify the longest practical value for max-age