// Copyright 2020 The Nakama Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package console

//build:ignore
//go:generate sh -c "protoc -I. -I../vendor -I../build/grpc-gateway-2.0.0-beta.5/third_party/googleapis -I../vendor/github.com/grpc-ecosystem/grpc-gateway -I../vendor/ --go_out=. --go_opt=plugins=grpc --go_opt=paths=source_relative --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=generate_unbound_methods=true --openapiv2_out=. --openapiv2_opt=logtostderr=true console.proto && protoc -I. -I../vendor -I../build/grpc-gateway-2.0.0-beta.5/third_party/googleapis -I../vendor/github.com/grpc-ecosystem/grpc-gateway --angular_out=filename=console.service.ts,service_name=ConsoleService:. console.proto && mv console.service.ts ui/src/app"