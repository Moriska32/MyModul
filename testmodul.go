package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/Moriska32/MyModul"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	wwwstop_times := "http://10.68.1.57/mgt/file?name=STOP_TIMES&type=TEXT"
	text = MGTModul.respons(wwwstop_times)

}
