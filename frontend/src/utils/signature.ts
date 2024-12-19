/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/
import hmacSHA256 from "crypto-js/hmac-sha256";

const GenerateSignature = (key: string, data: string) => {
  const hmac = hmacSHA256(data, key);
  return hmac.toString();
};

export default GenerateSignature;
