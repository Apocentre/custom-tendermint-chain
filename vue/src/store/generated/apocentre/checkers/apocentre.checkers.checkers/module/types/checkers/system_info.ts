/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "apocentre.checkers.checkers";

export interface SystemInfo {
  nextId: number;
}

const baseSystemInfo: object = { nextId: 0 };

export const SystemInfo = {
  encode(message: SystemInfo, writer: Writer = Writer.create()): Writer {
    if (message.nextId !== 0) {
      writer.uint32(8).int32(message.nextId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): SystemInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSystemInfo } as SystemInfo;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nextId = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SystemInfo {
    const message = { ...baseSystemInfo } as SystemInfo;
    if (object.nextId !== undefined && object.nextId !== null) {
      message.nextId = Number(object.nextId);
    } else {
      message.nextId = 0;
    }
    return message;
  },

  toJSON(message: SystemInfo): unknown {
    const obj: any = {};
    message.nextId !== undefined && (obj.nextId = message.nextId);
    return obj;
  },

  fromPartial(object: DeepPartial<SystemInfo>): SystemInfo {
    const message = { ...baseSystemInfo } as SystemInfo;
    if (object.nextId !== undefined && object.nextId !== null) {
      message.nextId = object.nextId;
    } else {
      message.nextId = 0;
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
