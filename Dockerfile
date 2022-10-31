FROM python:3.11-alpine

ENV PYTHONUSERBASE=/usr/src/app
ENV PATH=$PYTHONUSERBASE/bin:$PATH

# By default the app listens on localhost:8080. This can be overridden by setting the HOST and PORT environment variables
EXPOSE 8080

RUN mkdir -p "$PYTHONUSERBASE"

# Create python user and group, and set permissions on directories
RUN addgroup -S app_user && \
    adduser -S app_user -G app_user -h "$PYTHONUSERBASE" && \
    chown -R app_user:root "$PYTHONUSERBASE"

# Set user home directory as default
WORKDIR $PYTHONUSERBASE

# Copy app to image
COPY python/* "$PYTHONUSERBASE"

# Install app requirements
RUN pip install --user --no-cache-dir -r "$PYTHONUSERBASE/requirements.txt"

# Set default startup user
USER app_user

# Run script on startup
CMD [ "python", "time_app.py" ]
