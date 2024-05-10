import React from 'react'
import { Input } from 'antd';
import './index.css'

const { TextArea } = Input;

function InputBox({handleSend,value, setValue}) {




  return (
    <div className="custom-search-bar">
    <TextArea  value={value}
        onChange={(e) => setValue(e.target.value)}
        autoSize={{ minRows: 1, maxRows: 20 }}className={"searchbox"} placeholder="How can I help build your team" />
    <svg xmlns="http://www.w3.org/2000/svg"  onClick={()=>handleSend(value)} fill="none" viewBox="0 0 24 24"  height="30" width="30" stroke-width="1.5" stroke="currentColor" aria-hidden="true" className="svgClass"><path stroke-linecap="round" stroke-linejoin="round" d="M6 12L3.269 3.126A59.768 59.768 0 0121.485 12 59.77 59.77 0 013.27 20.876L5.999 12zm0 0h7.5"></path></svg>
    </div>
  )
}

export default InputBox