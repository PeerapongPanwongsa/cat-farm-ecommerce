// api.js
import axios from 'axios';

const API_URL = 'http://localhost:8080';

export const fetchCats = async () => {
  try {
    const response = await axios.get(`${API_URL}/api/cats`);
    return response.data;
  } catch (error) {
    console.error('Error fetching cats:', error);
    return [];
  }
};

// สามารถเพิ่มฟังก์ชันสำหรับเรียก API อื่นๆ ที่นี่ได้ เช่น
// export const fetchCatById = async (id) => { ... };