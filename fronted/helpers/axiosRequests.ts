import { useRouter, NextRouter } from 'next/dist/client/router'


const axios = require('axios');


export const postSignup = (email: string, password: string, router: NextRouter) => {
    axios.post(
        'http://localhost:8080/user/signup',
        { email: email, password: password }
    )
    .then((resp: any) => {
        postSignin(email, password, router)
    })
    .catch((e: any) => {
        if (e.response)
            alert(e.response.data)
        else
            console.log(e);
    });
}


export const postSignin = (email: string, password: string, router: NextRouter) => {
    axios.post(
        'http://localhost:8080/user/signin',
        { email: email, password: password },
    )
        .then((r: any) => {
            localStorage.setItem('token', r.data);
            window.dispatchEvent(new Event('storage'))
            getLists(router);
            router.push('/lists');
        })
        .catch((e: any) => {
            if (e.response)
                alert(e.response.data)
            else
                console.log(e);
        })
}

export const getLists = (router: NextRouter) => {
    let token = localStorage.token
    axios.get('http://localhost:8080/todo/lists', {headers: { Authorization: `Bearer ${token}` }}).then((r: any) => {
        console.log(r.data)
    })
}

export const postLists = (listName: string, router: NextRouter) => {
    axios.post('http://localhost:8080/todo/lists', {
            name: listName
        }, {
            headers: {
                'Authorization': `Bearer ${localStorage.token}`
            }
        })
        .then((r: any) => {
            getLists(router)
        })
        .catch((e: any) => {
            if (e.response) {
                if (e.response.status == 401)
                    router.push('/signin')
                alert(e.response.data)
            }
            else
                console.log(e);
        });
}

export const deleteLists = (id: number, router: NextRouter) => {
    axios.delete(`http://localhost:8080/todo/lists/${id}`, {
            headers: {
                'Authorization': `Bearer ${localStorage.token}`
            }
        })
        .then((r: any) => {
            getLists(router)
        })
        .catch((e: any) => {
            if (e.response) {
                if (e.response.status == 401)
                    router.push('/signin')
                alert(e.response.data)
            }
            else
                console.log(e);
        });
}

export const putLists = (id: number, name: string, router: NextRouter) => {
    axios.put(`http://localhost:8080/todo/lists/${id}`, {
            name: name
        }, {
            headers: {
                'Authorization': `Bearer ${localStorage.token}`
            },

        })
        .then((r: any) => {
            getLists(router)
        })
        .catch((e: any) => {
            if (e.response) {
                if (e.response.status == 401)
                    router.push('/signin')
                alert(e.response.data)
            }
            else
                console.log(e);
        });
}

export const getTasks = (list_id: number, router: NextRouter) => {
    axios.get(`http://localhost:8080/todo/lists/${list_id}/tasks`, {
            headers: {
                'Authorization': `Bearer ${localStorage.token}`
            }
        })
        .then((r: any) => {
            localStorage.setItem('tasksForCurrentList', JSON.stringify(r.data))
            window.dispatchEvent(new Event('storage'))
        })
        .catch((e: any) => {
            if (e.response) {
                if (e.response.status == 401)
                    router.push('/signin')
                alert(e.response.data)
            }
            else
                console.log(e);
        });
}

export const postTasks = (list_id: number, taskName: string, router: NextRouter) => {
    axios.post(`http://localhost:8080/todo/lists/${list_id}/tasks`, {
            'task_name': taskName,
            'description': '',
            'status': 'open'
        }, {
            headers: {
                'Authorization': `Bearer ${localStorage.token}`
            }
        })
        .then((r: any) => {
            getTasks(list_id, router)
        })
        .catch((e: any) => {
            if (e.response) {
                if (e.response.status == 401)
                    router.push('/signin')
                alert(e.response.data)
            }
            else
                console.log(e);
        });
}

export const deleteTasks = (list_id: number, task_id: number, router: NextRouter) => {
    axios.delete(`http://localhost:8080/todo/lists/${list_id}/tasks/${task_id}`, {
            headers: {
                'Authorization': `Bearer ${localStorage.token}`
            }
        })
        .then((r: any) => {
            getTasks(list_id, router)
        })
        .catch((e: any) => {
            if (e.response) {
                if (e.response.status == 401)
                    router.push('/signin')
                alert(e.response.data)
            }
            else
                console.log(e);
        });
}

export const putTasks = (list_id: number, task_id: number, name: string, description: string, status: string, router: NextRouter) => {
    axios.put(`http://localhost:8080/todo/lists/${list_id}/tasks/${task_id}`, {
            'task_name': name,
            'description': description,
            'status': status

        }, {
            headers: {
                'Authorization': `Bearer ${localStorage.token}`
            },

        })
        .then((r: any) => {
            getTasks(list_id, router)
        })
        .catch((e: any) => {
            if (e.response) {
                if (e.response.status == 401)
                    router.push('/signin')
                alert(e.response.data)
            }
            else
                console.log(e);
        });
}
