function getItemFromLocalStorage(key) {
    return new Promise((resolve, reject) => {
      try {
        const item = localStorage.getItem(key);
        resolve(item);
      } catch (error) {
        reject(error);
      }
    });
  }



  async function setUserData() {
        
    user = await getItemFromLocalStorage('user');
    alert(user)
    try{
        user = JSON.parse(user)
 
    }
    catch (ex){
        alert(ex)
    }
    try {
     
        const response = await fetch('/api/stuent_data?s_name=' + encodeURIComponent(user.uName)+'&s_password='+encodeURIComponent(user.uDate), {
        method: 'GET',
    });
    if (!response.ok) {
        throw new Error('Ошибка запроса: ' + response.statusText);
    }


    const data = await response.json();
    alert(data)
    setUserCardGroup(data)
setUserCardName(user.uName)
    return JSON.stringify(data, null, 2);
    } catch (error) {
        console.error('Ошибка:', error);
        alert('Ошибка: ' + error.message);
    }
}

function setUserCardGroup(text) {
  document.getElementById('userCardGroup').textContent = text;
}

function setUserCardName(text) {
  document.getElementById('userCardName').textContent = text;
}