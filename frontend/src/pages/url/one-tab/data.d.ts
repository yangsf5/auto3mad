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
