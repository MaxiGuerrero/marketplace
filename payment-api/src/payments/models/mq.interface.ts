interface IMqConnector {
  sendMessage<T extends object>(data: T): Promise<void>;
}

export { IMqConnector };
