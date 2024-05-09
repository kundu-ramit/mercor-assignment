import React from 'react';
import marcusLogo from '../../images/mercor_logo.png'; 
import { Image } from 'antd';

function Intro() {
  return (
    <div style={{ paddingTop: "30px", textAlign: "center" }}>
      <Image src={marcusLogo} />
      <p style={{ fontWeight: "bold", fontSize: "25px" }}>Marcus</p>
      <p style={{ fontSize: "20px" }}>AI Assistant at Mercor</p>
      <p style={{ fontSize: "20px", color: "#666" }}>Made with ❤️ in San Francisco</p>
    </div>
  );
}

export default Intro;
