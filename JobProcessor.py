from multiprocessing import (
  Pool,
  Process,
)

import subprocess, sys

import socket
import os

pools = {}


def log(text):
  print(data_rcvd)
  sys.stdout.flush()

#####################

def run_command(command):
  process = subprocess.Popen(
    command,
    shell=True,
    stdout=subprocess.PIPE,
    stderr=subprocess.STDOUT,
  )

  pid = process.pid

  stdout_lines = []

  for line in process.stdout:
    stdout_lines.append(line.rstrip())
    log(line)

  process.wait()

  output, error = process.communicate()
  returncode = process.returncode

  log(pid)
  log(returncode)
  log(stdout_lines)
  log(output)



def test_run_command():
  command = 'sleep 1;echo hey;sleep 4;echo foo;sleep 4;echo bar'
  p1 = Process(name='P1', target=run_command, args=(command,))
  p2 = Process(name='P2', target=run_command, args=(command,))

  p1.start()
  p2.start()

  # p1.is_alive()
  # p1.terminate()
  # p1.pid
  # p1.exitcode

  p1.join()
  p2.join()
#####################

# should be threaded looping, so that it may fire subprocess commands and interface with database
def start_socket_server(sock_path):
  os.unlink(sock_path) # if already exists
  log('Waiting for connection at {} ...'.format(sock_path))
  sock = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
  sock.bind(sock_path)
  sock.listen(1)
  while True:
    connection, client_address = sock.accept()
    data_rcvd = []
    while 1:
      data = connection.recv(1024)
      if data:
        data_rcvd.append(data)
      else:
        break
    log(data_rcvd)
    connection.close()

    if b'exit' in b''.join(data_rcvd).lower():
      # echo 'exit' | socat - UNIX-CONNECT:test.sock
      break

def test_socket_server():
  start_socket_server('test.sock')