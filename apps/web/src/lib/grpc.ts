import { createClient, type Transport } from "@connectrpc/connect";
import { createGrpcWebTransport } from "@connectrpc/connect-web";
import { EchoService } from "@repo/protobuf/ts/proto/echo/echo_pb.ts";
import { GameService } from "@repo/protobuf/ts/proto/game/game_pb.ts";
import { UserService } from "@repo/protobuf/ts/proto/user/user_pb.ts";

declare global {
  interface BigInt {
    toJSON(): string;
  }
}

BigInt.prototype.toJSON = function () {
  return this.toString();
};

const apiUrl = "http://localhost:3000";

export const transport: Transport = createGrpcWebTransport({
  baseUrl: apiUrl,
  fetch: (input, init) =>
    fetch(input, {
      ...init,
      credentials: "include",
    }),
});

export const GrpcServices = [EchoService, GameService, UserService];
export const GrpcMethods = GrpcServices.flatMap((s) => s.methods);

export const echoClient = createClient(EchoService, transport);
export const gameClient = createClient(GameService, transport);
export const userClient = createClient(UserService, transport);
