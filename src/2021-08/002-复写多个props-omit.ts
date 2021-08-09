interface AppLoader {
  name: string;
  size: number;
  api: () => Promise<any>;
}

interface AppDownLoader extends Omit<AppLoader, "api" | "size"> {
  api: string;
  size: () => number;
}
