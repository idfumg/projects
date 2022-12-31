import http from 'k6/http';

export let options = {
    vus: 10,
    duration: '10s'
}

export default function() {
    // http.get('http://localhost:8080/products')
    // http.get('http://host.docker.internal:8080/products')
}