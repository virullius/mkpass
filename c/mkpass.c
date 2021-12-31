#include <stdio.h>
#include <string.h>
#include <stdbool.h>
#include <stdlib.h>
#include <sys/time.h>

const char VERSION[] = "0.1.0";
const int MAX_TRIES = 10;

const char UPPER[] = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
const char LOWER[] = "abcdefghijklmnopqrstuvwxyz";
const char NUMBER[] = "0123456789";
const char SYMBOL[] = "`~!@#$%^&*()_+-=[]\\{}|;':\",./<>?";

struct options {
  short upper;
  short lower;
  short number;
  short symbol;
  int length;
};

struct options default_options = {
  .upper = true,
  .lower = true,
  .number = true,
  .symbol = true,
  .length = 16
};

struct generator {
  struct options opts;
  char *charset;
};

struct generator new_generator(struct options opts) {
  struct timeval t;
  gettimeofday(&t, NULL);
  srand(t.tv_sec + t.tv_usec);
  struct generator gen;
  gen.opts = opts;
  char *cs = NULL;
  if (opts.upper) {
    int sz = 0;
    if (cs != NULL) {
      sz = strlen(cs);
    }
    void *newp = realloc(cs, sz + strlen(UPPER));
    if (newp == NULL) {
      fprintf(stderr, "Failed to allocate memory for upper\n");
      exit(1);
    }
    cs = newp;
    strcat(cs, UPPER);
  }
  if (opts.lower) {
    int sz = 0;
    if (cs != NULL) {
      sz = strlen(cs);
    }
    void *newp = realloc(cs, sz + strlen(LOWER));
    if (newp == NULL) {
      fprintf(stderr, "Failed to allocate memory for lower\n");
      exit(1);
    }
    cs = newp;
    strcat(cs, LOWER);
  }
  if (opts.number) {
    int sz = 0;
    if (cs != NULL) {
      sz = strlen(cs);
    }
    void *newp = realloc(cs, sz + strlen(NUMBER));
    if (newp == NULL) {
      fprintf(stderr, "Failed to allocate memory for number\n");
      exit(1);
    }
    cs = newp;
    strcat(cs, NUMBER);
  }
  if (opts.symbol) {
    int sz = 0;
    if (cs != NULL) {
      sz = strlen(cs);
    }
    void *newp = realloc(cs, sz + strlen(SYMBOL));
    if (newp == NULL) {
      fprintf(stderr, "Failed to allocate memory for symbol\n");
      exit(1);
    }
    cs = newp;
    strcat(cs, SYMBOL);
  }

  gen.charset = cs;

  return gen;
}

int generate(struct generator gen, char **str) {
  char *np = calloc(sizeof(char), gen.opts.length + 1);
  if (np == NULL) {
    fprintf(stderr, "Failed to allocate memory for generated string\n");
    return 0;
  }

  for (int i = 0; i < gen.opts.length; i++) {
    int x = rand() % strlen(gen.charset);
    np[i] = (char)gen.charset[x];
  }
  *str = np;

  return 1;
}

int main(int argc, char *argv[]) {

  struct generator gen = new_generator(default_options);

  char *pw = NULL;
  if (!generate(gen, &pw)) {
    fprintf(stderr, "Failed to generate string\n");
    return 0;
  }
  printf("%s\n", pw);

  return 1;
}
