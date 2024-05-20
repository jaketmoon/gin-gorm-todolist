document.getElementById('loginForm').addEventListener('submit', function(e) {
    e.preventDefault();

    // 获取用户输入的用户名和密码
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    // 发送登录请求到后端
    fetch('http://localhost:8080/User/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        // 注意，这里我们将对象转换为x-www-form-urlencoded格式
        body: `name=${encodeURIComponent(username)}&password=${encodeURIComponent(password)}`
    })
        .then(response => {
            // 通过响应状态码判断请求是否成功
            if(response.ok) {
                return response.json();
            }
            throw new Error('Something went wrong');
        })
        .then(data => {
            // 登录成功的操作
            console.log('Login Successful:', data);
            // 保存token到localStorage 或 sessionStorge
            localStorage.setItem('token', data.token);
            // 跳转到待办事项页面
            window.location.href = 'todo.html';
        })
        .catch(error => {
            // 登录失败的操作
            console.error('Login Failed:', error);
            alert('Login failed');
        });
});