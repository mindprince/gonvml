/*
Copyright 2017 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

#include <stddef.h>
#include <dlfcn.h>

#include "nvml.h"

// nvmlHandle is the handle for dynamically loaded libnvidia-ml.so
void *nvmlHandle;

/**
 * Loads the "libnvidia-ml.so.1" shared library and initializes NVML.
 * Call this before calling any other methods.
 */
nvmlReturn_t nvmlInit_dl(void) {
  nvmlHandle = dlopen("libnvidia-ml.so.1", RTLD_LAZY | RTLD_GLOBAL);
  if (nvmlHandle == NULL) {
    return (NVML_ERROR_LIBRARY_NOT_FOUND);
  }
  return (nvmlInit());
}

/**
 * Shuts down NVML and decrements the reference count on the dynamically loaded
 * "libnvidia-ml.so.1" library.
 * Call this once NVML is no longer being used.
 */
nvmlReturn_t nvmlShutdown_dl(void) {
  if (nvmlHandle == NULL) {
    return NVML_SUCCESS;
  }
  nvmlReturn_t r = nvmlShutdown();
  if (r != NVML_SUCCESS) {
    return (r);
  }
  return (dlclose(nvmlHandle) ? NVML_ERROR_UNKNOWN : NVML_SUCCESS);
}
