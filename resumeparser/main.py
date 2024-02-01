import functions_framework


@functions_framework.cloud_event
def parse_resume(cloud_event):
    print(cloud_event.data)
    
