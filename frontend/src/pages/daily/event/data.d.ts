export type RoutineInfo = {
  id: number,
  icon: string,
  short_name: string,
  event_scope: string,
  will_spend: number,
  history_spend: number,
  today_spend: number,
  week_will_spend: number,
  week_spend: number,
  month_will_spend: number,
  month_spend: number,
  total_will_spend: number,
  total_spend: number,
  week_passed: number,
  object: number,
  object_unit: string,
  progress: number,
  start_date: string,
};

export type EventInfo = {
  id: number,
  date: string,
  start_time: string,
  end_time: string,
  routine_id: number,
  spend: number,
};

export type EventAPI = {
  events: EventInfo[],
};

export type EditEventInfo = {
  id: number,
  date: string,
  start_time: string,
  end_time: string,
  specific_event: string,
  routine_id: number,
};
