document.getElementById('uploadForm').addEventListener('submit', function(event) {
    event.preventDefault();

    let fileInput = document.getElementById('uploadFile');
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
                // If the response status is not in the range 200-299
                throw new Error('File upload failed'); // Throw an error
            }
            // If the response status is in the range 200-299
            return response.text(); // Return the response body as text
        })
        .then(data => {
            // Handle the successful response here
            document.getElementById('uploadMessage').innerText = data;
            fileInput.value = ''; // Clear the input
        })
        .catch(error => {
            // Handle any errors that occurred during the fetch
            console.error('Error:', error);
            document.getElementById('uploadMessage').innerText = 'File upload failed';
        });

});

document.getElementById('downloadForm').addEventListener('submit', function(event) {
    event.preventDefault();

    let filename = document.getElementById('downloadFileName').value;

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
});
