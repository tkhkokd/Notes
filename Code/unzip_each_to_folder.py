
# This code will unzip zipped folders to directories
# the directory name will be the same as the zip
# zip that only contains a single file will also be unzipped to a (new) directory

import os
import zipfile

for root, dirs, files in os.walk("."):
    # files => array of zip
    print(files)

    for f in files:
        zip_ref = zipfile.ZipFile(f, 'r')
        zip_ref.extractall(f[:-4])
        zip_ref.close()
