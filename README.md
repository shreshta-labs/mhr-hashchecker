### hashchecker - Check for file hashes in Team Cymru's Malware Hash Registry.

hashchecker.go is a go program which will generate a hash for files in a directory, recursively. It will then check for the hash on Team Cymru's Malware Hash Registry (https://www.team-cymru.com/mhr).

The bad hashes are then stored along with absolute path of the file in a file for the user to review.

hashchecker.go can be configured as a cronjob/Scheduled task such that it can be automated on a daily basis.


Before executing, modify the following in hashchecker

1. /path/to/badhashes.txt - The absolute path of the text file which will contain the bad hashes matched in the Team Cymru's database

2. /path/of/directory - The absolute path of the directory to scan

3. /path/to/store/hashes.txt - The absolute path of text file where the hash and absolute path of the files will be stored


git clone https://github.com/shreshta-labs/mhr-hashchecker.git

cd team-cymru-mhr

go run hashchecker.go


For any bugs, suggestions, please send an email to mhr-hashchecker@shreshtait.com
