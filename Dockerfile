FROM python:latest

RUN pip install --no-cache-dir jupyterlab

WORKDIR /work

CMD ["jupyter", "lab", "--ip=0.0.0.0", "--allow-root", "--no-browser"]
