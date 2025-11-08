import axios, { AxiosRequestConfig, AxiosResponse } from "axios";

export class ApiClient {
  private baseUrl: string;

  constructor(baseUrl: string) {
    this.baseUrl = baseUrl;
  }

  // General GET request
  async get<T>(endpoint: string, config?: AxiosRequestConfig): Promise<T> {
    try {
      const response: AxiosResponse<T> = await axios.get(
        `${this.baseUrl}${endpoint}`,
        config
      );
      return response.data;
    } catch (error) {
      throw this.handleError(error);
    }
  }

  // General POST request
  async post<T>(
    endpoint: string,
    data: any,
    config?: AxiosRequestConfig
  ): Promise<T> {
    try {
      const response: AxiosResponse<T> = await axios.post(
        `${this.baseUrl}${endpoint}`,
        data,
        config
      );
      return response.data;
    } catch (error) {
      throw this.handleError(error);
    }
  }

  // General PUT request
  async put<T>(
    endpoint: string,
    data: any,
    config?: AxiosRequestConfig
  ): Promise<T> {
    try {
      const response: AxiosResponse<T> = await axios.put(
        `${this.baseUrl}${endpoint}`,
        data,
        config
      );
      return response.data;
    } catch (error) {
      throw this.handleError(error);
    }
  }

  // General DELETE request
  async delete<T>(endpoint: string, config?: AxiosRequestConfig): Promise<T> {
    try {
      const response: AxiosResponse<T> = await axios.delete(
        `${this.baseUrl}${endpoint}`,
        config
      );
      return response.data;
    } catch (error) {
      throw this.handleError(error);
    }
  }

  // Error handling
  private handleError(error: any): Error {
    if (error.response) {
      // Server error
      return new Error(
        `Error: ${error.response.status} - ${error.response.data}`
      );
    } else if (error.request) {
      // No response from server
      return new Error("No response received from server.");
    } else {
      // Error setting up the request
      return new Error(`Request setup error: ${error.message}`);
    }
  }
}
