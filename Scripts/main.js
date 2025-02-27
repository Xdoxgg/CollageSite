async function getUserData() {
    const user = localStorage.getItem('user')

    user = JSON.parse(user)
    return loadDataIn(user.name , user.password)

}


async function loadDataIn(name, password){

    event.preventDefault(); // Предотвращаем перезагрузку страницы

    try {
     // Отправляем запрос на сервер
    const response = await fetch('/api/student_by_input_data?s_name=' + encodeURIComponent(name)+'&s_password='+encodeURIComponent(password), {
        method: 'GET',
    });
    if (!response.ok) {
        throw new Error('Ошибка запроса: ' + response.statusText);
    }

    // Получаем данные из ответа
    const data = await response.json();
  

    // Отображаем результат
    return data;
    } catch (error) {
    console.error('Ошибка:', error);
    alert('Ошибка: ' + error.message);
}

}