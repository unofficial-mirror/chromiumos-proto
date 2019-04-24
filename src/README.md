# proto/src

This folder contains the protobuf source files for the Chromite Build API and Recipes implementations.

## chromite/api

This is the core API definitions, including all of the services and most of the messages.

The build_api.proto contains service and method options that define some of the implementation details. Specifically, the service options define the controller module that implements the methods, and optionally enforces whether or not the methods should run in the chroot. The method options allow overriding the default function name for the implementation, and overriding the service-level chroot setting.

* Service
  * module - Required
  * service_chroot_assert - Recommended
* Method
  * implementation_name - Optional
  * method_chroot_assert - Optional

These options are all implementation details that should not affect consumers of the endpoints, but are important details for anyone implementing endpoints.

## chromiumos

This folder contains more widely shared proto files.


## device


## testplans
