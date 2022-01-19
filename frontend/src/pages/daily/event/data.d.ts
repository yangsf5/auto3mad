export type RoutineInfo = {
  id: number,
  icon: string,
  short_name: string,
  event: string,
  will_spend: number,
  spend: number,
  total_spend: number,
};

export type EventInfo = {
  id: number,
  start_time: string,
  end_time: string,
  specific_event: string,
  routine_event: string,
  spend: number,
}