// 代办事项应用的前端代码
function handleResponse(response) {
    if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
    }
    return response.json();
}

// 添加代办事项
function addTodoItem() {
    const todoText = document.getElementById('todo-input').value;
    fetch('http://localhost:8080/Todo/', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ title: todoText, status: false }),
    })
        .then(handleResponse)
        .then(data => {
            console.log('Todo added:', data);
            document.getElementById('todo-input').value = '';
            fetchTodos(); // 刷新列表
        })
        .catch(console.error);
}

// 展示代办事项
function fetchTodos() {
    fetch('http://localhost:8080/Todo/')
        .then(response => response.json())
        .then(json => {
            const todosArray = json.data[1];
            const list = document.getElementById('todo-list');
            list.innerHTML = todosArray.map(todo => `
                <li id="todo-item-${todo.id}">
                    <span>${todo.title}</span>
                    <span class="status">${todo.status ? '-    已完成' : '-    未完成'}</span>
                    <button onclick="deleteTodoItem(${todo.id})">删除</button>
                    <button onclick="toggleComplete(${todo.id}, ${todo.status})">
                        ${todo.status ? '撤销完成' : '标记完成'}
                    </button>
                </li>
            `).join('');
        })
        .catch(error => {
            console.error('Error fetching todos:', error);
        });
}

// You should also have a handleResponse function or remove the call to it in the toggleComplete function
// Assuming handleResponse looks like this, you can adjust as needed:
function handleResponse(response) {
    if (!response.ok) {
        throw new Error('Network response was not ok ' + response.statusText);
    }
    return response.json();
}

function toggleComplete(id, completed) {
    fetch(`http://localhost:8080/Todo/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ status: !completed }), // Toggle the completion status
    })
        .then(response => response.json())
        .then(json => {
            const updatedTodo = json.data[1]; // Assume the updated todo is returned at this index
            console.log('Todo updated:', updatedTodo);
            // Now, update the display based on the new completion status
            document.querySelector(`#todo-item-${id} .status`).textContent = updatedTodo.status ? 'Completed' : 'Not Completed';
            fetchTodos();
        })
        .catch(console.error);
}

// 删除代办事项
function deleteTodoItem(id) {
    fetch(`http://localhost:8080/Todo/${id}`, {
        method: 'DELETE',
    })
        .then(handleResponse)
        .then(data => {
            console.log('Todo deleted:', data);
            fetchTodos(); // Refresh the list
        })
        .catch(console.error);
}

// 初始化
document.addEventListener('DOMContentLoaded', fetchTodos);