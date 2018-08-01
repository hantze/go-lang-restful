# Copyright 2016 The Kubernetes Authors.
# Modifications Copyright 2018 Payfazz.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM ARG_FROM

LABEL maintainer Ricky Winata <ricky@payfazz.com>

# in order to support https request
RUN apk add --no-cache ca-certificates

ADD bin/ARG_ARCH/ARG_BIN /ARG_BIN

USER nobody:nobody
ENTRYPOINT ["/ARG_BIN"]