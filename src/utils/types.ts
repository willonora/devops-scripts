// types.ts

import * as path from 'path';
import { EOL } from 'os';

export type Config = {
  source: string;
  destination: string;
  command: string;
  options?: {
    [key: string]: string | number;
  };
};

export type ShellCommand = {
  command: string;
  options: { [key: string]: string | number };
};

export type File = {
  name: string;
  path: string;
  contents: string;
};

export type Directory = {
  name: string;
  path: string;
  files: File[];
  subdirectories: Directory[];
};

export type DevOpsScript = {
  name: string;
  description: string;
  config: Config;
  commands: ShellCommand[];
};

export type DevOpsScriptConfig = {
  sourceDir: string;
  destinationDir: string;
  command: string;
  options?: {
    [key: string]: string | number;
  };
};