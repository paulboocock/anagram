FROM alpine:3.14

ADD build/output/anagram_linux_amd64 /root/anagram_linux_amd64

ENTRYPOINT ["./root/anagram_linux_amd64"]
