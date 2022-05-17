export type StatChatItem = {
  routine: string,
  month: string,
  spend: number,
};

export type StatTableItem = {
  routine: string,
  spends: number[],
};

export type StatRet = {
  chart: StatChatItem[],
  table: StatTableItem[],
  months: string[],
};
