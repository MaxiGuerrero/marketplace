import axios from 'axios';
import { NextFunction, Request, Response } from 'express';
import jwtDecode from 'jwt-decode';
import { config, User } from '@shared/index';
import logger from '@shared/utils/logger';

async function authMiddleware(req: Request, res: Response, next: NextFunction) {
  try {
    const token = req.headers.authorization;
    if (!token) {
      return res.status(401).json({ error: 'Access token is missing or is invalid' });
    }
    const result = await axios.post(`${config.SECURITY_API_URL}/token/validate`, null, {
      headers: {
        Authorization: `${token}`,
      },
      timeout: 5000,
    });
    if (result.status !== 200) {
      return res.status(401).json({ error: 'Access token is missing or is invalid' });
    }
    req.body.user = jwtDecode<User>(token);
    logger.debug(`Security api response with status: ${result.status}`);
    return next();
  } catch (error: unknown) {
    if (axios.isAxiosError(error)) {
      logger.debug(`Security api response with status: ${error.response?.status}: ${error.response?.data}`);
      return res.status(401).json({ error: 'Access token is missing or is invalid' });
    }
    logger.error(error);
    throw error;
  }
}

export { authMiddleware };
