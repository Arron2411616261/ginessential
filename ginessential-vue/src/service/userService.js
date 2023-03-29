import request from '@/utils/request';

const register = ({ name, telephone, password }) => {
  return request.post('auth/register', { name, telephone, password });
};

const info = () => {
  return request.get('auth/info');
};

export default {
  register,
  info,
};
