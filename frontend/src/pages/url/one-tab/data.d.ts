export type UrlItem = {
  icon: string,
  url: string,
  title: string,
};

export type UrlGroup = {
  title: string;
  icon: string;
  urls: UrlItem[];
};

export type GroupInfo = {
  id: number,
  title: string,
  icon: string,
};
