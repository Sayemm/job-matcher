import axios from 'axios';

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api';

export const uploadResume = async (file) => {
  const formData = new FormData();
  formData.append('resume', file);

  const response = await axios.post(
    `${API_BASE_URL}/resume/upload-and-match`,
    formData,
    {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    }
  );

  return response.data;
};

export const getJobsByCluster = async (clusterId, page = 1, pageSize = 20) => {
  const response = await axios.get(
    `${API_BASE_URL}/jobs/cluster/${clusterId}`,
    {
      params: { page, page_size: pageSize },
    }
  );

  return response.data;
};

export const getAllJobs = async (page = 1, pageSize = 20) => {
  const response = await axios.get(`${API_BASE_URL}/jobs`, {
    params: { page, page_size: pageSize },
  });

  return response.data;
};

export const getJobById = async (jobId) => {
  const response = await axios.get(`${API_BASE_URL}/jobs/${jobId}`);
  return response.data;
};