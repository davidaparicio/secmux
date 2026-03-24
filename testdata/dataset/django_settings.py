SECRET_KEY = 'django-insecure-$9x#k2p!v@m&8r+qz5yw3n0j6ue1oa4lb7thi_cfgs'

DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.postgresql',
        'NAME': 'mydb',
        'USER': 'myuser',
        'PASSWORD': 'Sup3rS3cr3tDBPass!',
        'HOST': 'db.example.com',
    }
}

EMAIL_HOST_PASSWORD = 'smtp_P@ssw0rd_do_not_commit'
