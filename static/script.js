const container = document.querySelector('.container');
const eventSource = new EventSource('http://localhost:6767/events');

eventSource.onmessage = function (event) {
  const text = event.data;
  const bubble = document.createElement('div');
  bubble.className = 'bubble';
  bubble.textContent = text;
  container.appendChild(bubble);

  setTimeout(function () {
    bubble.style.opacity = 1;
    bubble.style.transform = 'translateY(-50px)';
    setTimeout(function () {
      bubble.style.opacity = 0;
      bubble.style.transform = 'translateY(-100px)';
      container.removeChild(bubble); // Remove the bubble from the container
    }, 1500);
  }, 0);
};
