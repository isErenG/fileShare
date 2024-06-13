document.addEventListener('DOMContentLoaded', init);


function init() {
    let uploadInput = document.getElementById("uploadButton");
    uploadInput.addEventListener('change', function(event) {
        event.preventDefault();
        upload();
    });
    console.log("yod");
    let downloadForm = document.getElementById("downloadForm");
    downloadForm.addEventListener('submit', function(event) {
        event.preventDefault();
        download();
    });
}

function upload() {
    let fileInput = document.getElementById('uploadButton');
    let file = fileInput.files[0];
    let formData = new FormData();
    formData.append('file', file);

    fetch('http://localhost:8080/upload', {
        method: 'POST',
        body: formData
    })
        .then(response => {
            if (!response.ok) {
                console.log(response.status);
                throw new Error('File upload failed');
            }
            return response.json(); // Expecting a JSON response
        })
        .then(data => {
            // Handle the successful response here
            document.getElementById('uploadMessage').innerText = `File uploaded successfully: ${data.file_code}`;
            fileInput.value = ''; // Clear the input
        })
        .catch(error => {
            // Handle any errors that occurred during the fetch
            console.error('Error:', error);
            document.getElementById('uploadMessage').innerText = 'File upload failed';
        });
}

function download() {
    let filename = document.getElementById('codeInput').value;
    console.log(filename);
    fetch(`http://localhost:8080/download?filename=${filename}`)
        .then(response => {
            if (!response.ok) {
                throw new Error('File not found');
            }
            return response.blob();
        })
        .then(blob => {
            let url = window.URL.createObjectURL(blob);
            let a = document.createElement('a');
            a.href = url;
            a.download = filename;
            document.body.appendChild(a); // Required for Firefox
            a.click();
            a.remove();
            document.getElementById('downloadMessage').innerText = 'File downloaded successfully';
        })
        .catch(error => {
            console.error('Error:', error);
            document.getElementById('downloadMessage').innerText = 'File download failed';
        });
}
