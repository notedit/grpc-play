import * as jspb from 'google-protobuf'



export class JoinRequest extends jspb.Message {
  getRoom(): string;
  setRoom(value: string): JoinRequest;

  getUser(): string;
  setUser(value: string): JoinRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): JoinRequest.AsObject;
  static toObject(includeInstance: boolean, msg: JoinRequest): JoinRequest.AsObject;
  static serializeBinaryToWriter(message: JoinRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): JoinRequest;
  static deserializeBinaryFromReader(message: JoinRequest, reader: jspb.BinaryReader): JoinRequest;
}

export namespace JoinRequest {
  export type AsObject = {
    room: string,
    user: string,
  }
}

export class JoinResponse extends jspb.Message {
  getCode(): number;
  setCode(value: number): JoinResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): JoinResponse.AsObject;
  static toObject(includeInstance: boolean, msg: JoinResponse): JoinResponse.AsObject;
  static serializeBinaryToWriter(message: JoinResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): JoinResponse;
  static deserializeBinaryFromReader(message: JoinResponse, reader: jspb.BinaryReader): JoinResponse;
}

export namespace JoinResponse {
  export type AsObject = {
    code: number,
  }
}

export class LeaveRequest extends jspb.Message {
  getRoom(): string;
  setRoom(value: string): LeaveRequest;

  getUser(): string;
  setUser(value: string): LeaveRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LeaveRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LeaveRequest): LeaveRequest.AsObject;
  static serializeBinaryToWriter(message: LeaveRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LeaveRequest;
  static deserializeBinaryFromReader(message: LeaveRequest, reader: jspb.BinaryReader): LeaveRequest;
}

export namespace LeaveRequest {
  export type AsObject = {
    room: string,
    user: string,
  }
}

export class LeaveResponse extends jspb.Message {
  getCode(): number;
  setCode(value: number): LeaveResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LeaveResponse.AsObject;
  static toObject(includeInstance: boolean, msg: LeaveResponse): LeaveResponse.AsObject;
  static serializeBinaryToWriter(message: LeaveResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LeaveResponse;
  static deserializeBinaryFromReader(message: LeaveResponse, reader: jspb.BinaryReader): LeaveResponse;
}

export namespace LeaveResponse {
  export type AsObject = {
    code: number,
  }
}

export class StreamRequest extends jspb.Message {
  getRoom(): string;
  setRoom(value: string): StreamRequest;

  getUser(): string;
  setUser(value: string): StreamRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StreamRequest.AsObject;
  static toObject(includeInstance: boolean, msg: StreamRequest): StreamRequest.AsObject;
  static serializeBinaryToWriter(message: StreamRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StreamRequest;
  static deserializeBinaryFromReader(message: StreamRequest, reader: jspb.BinaryReader): StreamRequest;
}

export namespace StreamRequest {
  export type AsObject = {
    room: string,
    user: string,
  }
}

export class Login extends jspb.Message {
  getUser(): string;
  setUser(value: string): Login;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Login.AsObject;
  static toObject(includeInstance: boolean, msg: Login): Login.AsObject;
  static serializeBinaryToWriter(message: Login, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Login;
  static deserializeBinaryFromReader(message: Login, reader: jspb.BinaryReader): Login;
}

export namespace Login {
  export type AsObject = {
    user: string,
  }
}

export class Leave extends jspb.Message {
  getUser(): string;
  setUser(value: string): Leave;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Leave.AsObject;
  static toObject(includeInstance: boolean, msg: Leave): Leave.AsObject;
  static serializeBinaryToWriter(message: Leave, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Leave;
  static deserializeBinaryFromReader(message: Leave, reader: jspb.BinaryReader): Leave;
}

export namespace Leave {
  export type AsObject = {
    user: string,
  }
}

export class Message extends jspb.Message {
  getUser(): string;
  setUser(value: string): Message;

  getMessage(): string;
  setMessage(value: string): Message;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Message.AsObject;
  static toObject(includeInstance: boolean, msg: Message): Message.AsObject;
  static serializeBinaryToWriter(message: Message, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Message;
  static deserializeBinaryFromReader(message: Message, reader: jspb.BinaryReader): Message;
}

export namespace Message {
  export type AsObject = {
    user: string,
    message: string,
  }
}

export class Shutdown extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Shutdown.AsObject;
  static toObject(includeInstance: boolean, msg: Shutdown): Shutdown.AsObject;
  static serializeBinaryToWriter(message: Shutdown, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Shutdown;
  static deserializeBinaryFromReader(message: Shutdown, reader: jspb.BinaryReader): Shutdown;
}

export namespace Shutdown {
  export type AsObject = {
  }
}

export class StreamResponse extends jspb.Message {
  getLogin(): Login | undefined;
  setLogin(value?: Login): StreamResponse;
  hasLogin(): boolean;
  clearLogin(): StreamResponse;

  getLeave(): Leave | undefined;
  setLeave(value?: Leave): StreamResponse;
  hasLeave(): boolean;
  clearLeave(): StreamResponse;

  getMessage(): Message | undefined;
  setMessage(value?: Message): StreamResponse;
  hasMessage(): boolean;
  clearMessage(): StreamResponse;

  getShutdown(): Shutdown | undefined;
  setShutdown(value?: Shutdown): StreamResponse;
  hasShutdown(): boolean;
  clearShutdown(): StreamResponse;

  getEventCase(): StreamResponse.EventCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StreamResponse.AsObject;
  static toObject(includeInstance: boolean, msg: StreamResponse): StreamResponse.AsObject;
  static serializeBinaryToWriter(message: StreamResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StreamResponse;
  static deserializeBinaryFromReader(message: StreamResponse, reader: jspb.BinaryReader): StreamResponse;
}

export namespace StreamResponse {
  export type AsObject = {
    login?: Login.AsObject,
    leave?: Leave.AsObject,
    message?: Message.AsObject,
    shutdown?: Shutdown.AsObject,
  }

  export enum EventCase { 
    EVENT_NOT_SET = 0,
    LOGIN = 1,
    LEAVE = 2,
    MESSAGE = 3,
    SHUTDOWN = 4,
  }
}

