# _Mandrill_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)
[![Build Status](https://travis-ci.com/omg-services/mandrill.svg?branch=master)](https://travis-ci.com/omg-services/mandrill)
[![codecov](https://codecov.io/gh/omg-services/mandrill/branch/master/graph/badge.svg)](https://codecov.io/gh/omg-services/mandrill)


An OMG service for Mandrill, is a powerful delivery service that can be used for personalized, one-to-one e-commerce emails, and automated transactional emails like password resets, order confirmations, and welcome messages.

## Usage in [Storyscript](https://storyscript.io/)

##### Send Email
```coffee
>>> mandrill send from:'abc@example.com' to:'xyz@example.com' message:'messageBody' templateName:'templateName' subject:'emailSubject'
[{"email": "receiverEmail","status": "emailStatus","_id": "operationId","reject_reason": "reasonToReject"}]
```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)

##### Send Email
```shell
$ omg run send -a from=<SENDER_MAIL_ADDRESS> -a to=<RECEIVER_EMAIL_ADDRESS> -a message=<EMAIL_MESSAGE_BODY> -a templateName=<TEMPLATE_NAME> -a subject=<EMAIL_SUBJECT> -e API_KEY=<API_KEY>
```
**Note**: The OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/heaptracetechnology/microservice-mandrill/blob/master/LICENSE).
