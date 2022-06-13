import subprocess

completed = subprocess.run(["ls", "-la"], capture_output=True, text=True)

completed1 = subprocess.run(["python3", "other.py"], capture_output=True, text=True)

print("args", completed.args)
print("returnCode", completed.returncode)
print("err", completed.stderr)
print("out", completed.stdout)

