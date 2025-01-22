import React from 'react'

function Logo({width = '100px'}) {
  return (
    <div style={{ width, height: 'auto', overflow: 'hidden' }}>
      <img 
        src="https://tse2.mm.bing.net/th?id=OIP.WFSVR9iaq6ADFxn8NFEYhgHaKQ&pid=Api&P=0&h=180" 
        alt="BLOG" 
        style={{ width: '40%', height: 'auto' }} 
      />
    </div>
  )
}
export default Logo