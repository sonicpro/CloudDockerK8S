FROM ubuntu
# Write the greeting.
RUN echo "привет мир!" > hello_world

# Create a script that prints the file contents to the terminal.
RUN echo "#!/bin/bash\n\
while true; do date; cat hello_world; sleep \$1; done" > print_loop.sh

# Set convenient working directory
WORKDIR /opt/helloworld

# Copy the file that contain the greeting into the working directory.
# ./ corresponds to the directory that is set by WORKDIR command.
RUN cp /hello_world ./
RUN cp /print_loop.sh ./

# Set execute bit for print_loop.sh
RUN chmod +x ./print_loop.sh

# CMD (or ENTRYPOINT) command runs on the container start.
# The command accepts shell command as its first argument. The rest of arguments are the shell command arguments.
CMD ["/opt/helloworld/print_loop.sh", "2"]
