import path from 'path';
import fs from 'fs';
import { Router } from 'express';

export default async function registerRoutes() {
  const routesDir = path.join(__dirname);
  const routeFiles = fs
    .readdirSync(routesDir)
    .filter((file) => (file.endsWith('.ts') || file.endsWith('.js')) && !file.includes('index'));
  const importedFiles = (await Promise.all(routeFiles.map((file) => import(path.join(routesDir, file))))) as {
    default: object;
  }[];
  const routes = importedFiles.map((object) => object.default as Router);
  return routes;
}
