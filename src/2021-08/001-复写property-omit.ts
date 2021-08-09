// https://stackoverflow.com/questions/41285211/overriding-interface-property-type-defined-in-typescript-d-ts-file

interface App {
  name: string;
  api: () => Promise<any>;
}

interface AppDown extends Omit<App, "api"> {
  api: string;
}

type App1 = {
  name: string;
  api: () => Promise<any>;
};

type AppDown1 = Omit<App1, "api"> & {
  api: string;
};
