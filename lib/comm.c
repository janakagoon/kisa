/*
 * comm.c:
 *	Communication routines "platform specific" for Raspberry Pi
 *	
 *	Copyright (c) 2016-2020 Sequent Microsystem
 *	<http://www.sequentmicrosystem.com>
 ***********************************************************************
 *	Author: Alexandru Burcea
 ***********************************************************************
 */
#include <stdio.h>
#include <stdint.h>
#include <string.h>
#include <unistd.h>
#include <fcntl.h>
#include <sys/ioctl.h>

#  if __has_include (<linux/i2c-dev.h>)
#    include <linux/i2c-dev.h>
#  endif

//#include <linux/i2c-dev.h>
#include "comm.h"

#define I2C_SLAVE	0x0703
#define I2C_SMBUS	0x0720	/* SMBus-level access */

#define I2C_SMBUS_READ	1
#define I2C_SMBUS_WRITE	0

// SMBus transaction types

#define I2C_SMBUS_QUICK		    0
#define I2C_SMBUS_BYTE		    1
#define I2C_SMBUS_BYTE_DATA	    2
#define I2C_SMBUS_WORD_DATA	    3
#define I2C_SMBUS_PROC_CALL	    4
#define I2C_SMBUS_BLOCK_DATA	    5
#define I2C_SMBUS_I2C_BLOCK_BROKEN  6
#define I2C_SMBUS_BLOCK_PROC_CALL   7		/* SMBus 2.0 */
#define I2C_SMBUS_I2C_BLOCK_DATA    8

// SMBus messages
#define I2C_SMBUS_BLOCK_MAX	32	/* As specified in SMBus standard */
#define I2C_SMBUS_I2C_BLOCK_MAX	32	/* Not specified but we use same structure */


int i2cSetup(int addr)
{
#ifdef DEBUG
	printf("comm.c: i2cSetup started. addr= %d\r\n", addr);
#endif

	int file;

    char filename[40];
    sprintf(filename, "/dev/i2c-1");

#ifdef DEBUG
    printf("comm.c: i2cSetup opening file. filename= %s\r\n", filename);
#endif

    if ( (file = open(filename, O_RDWR)) < 0)
    {
#ifdef DEBUG
        printf("comm.c: i2cSetup open(filename, O_RDWR) failed\r\n");
#endif
        return -1;
    }

    if (ioctl(file, I2C_SLAVE, addr) < 0)
    {
#ifdef DEBUG
        printf("comm.c: i2cSetup ioctl(file, I2C_SLAVE, addr) failed\r\n");
#endif
        return -1;
    }

#ifdef DEBUG
    printf("comm.c: i2cSetup success. file= %d\r\n", file);
#endif
	return file;
}

int i2cMem8Read(int dev, int add, uint8_t* buff, int size)
{

#ifdef DEBUG
	printf("comm.c: i2cMem8Read started. dev= %d, add= %d\r\n", dev, add);
#endif

	uint8_t intBuff[I2C_SMBUS_BLOCK_MAX];

	if (NULL == buff)
	{
#ifdef DEBUG
		printf("comm.c: i2cMem8Read error. NULL == buff\r\n");
#endif
		return -1;
	}

	if (size > I2C_SMBUS_BLOCK_MAX)
	{
#ifdef DEBUG
	    printf("comm.c: i2cMem8Read error. size > I2C_SMBUS_BLOCK_MAX\r\n");
#endif
		return -1;
	}

	intBuff[0] = 0xff & add;

	int wr = write(dev, intBuff, 1);

	if (wr != 1)
	{
#ifdef DEBUG
	    printf("comm.c: i2cMem8Read write(dev, intBuff, 1) != 1 error. write(...) = %d.\r\n", wr);
#endif
		return -1;
	}

	int rr = read(dev, buff, size);

	if (rr != size)
	{
#ifdef DEBUG
	    printf("comm.c: i2cMem8Read read(dev, buff, size) != size error. read(...)= %d, size= %d\r\n", rr, size);
#endif
		return -1;
	}

#ifdef DEBUG
	printf("comm.c: i2cMem8Read success\r\n", dev, add);
#endif

	return 0;
}

int i2cMem8Write(int dev, int add, uint8_t* buff, int size)
{
#ifdef DEBUG
	printf("comm.c: i2cMem8Write started. dev= %d, add= %d\r\n", dev, add);
#endif

	uint8_t intBuff[I2C_SMBUS_BLOCK_MAX];

	if (NULL == buff)
	{
#ifdef DEBUG
		printf("comm.c: i2cMem8Write error. NULL == buff\r\n");
#endif
		return -1;
	}

	if (size > I2C_SMBUS_BLOCK_MAX - 1)
	{
#ifdef DEBUG
	    printf("comm.c: i2cMem8Write error. size > I2C_SMBUS_BLOCK_MAX - 1\r\n");
#endif
		return -1;
	}

	intBuff[0] = 0xff & add;
	memcpy(&intBuff[1], buff, size);

	if (write(dev, intBuff, size + 1) != size + 1)
	{
#ifdef DEBUG
	    printf("comm.c: i2cMem8Write error. write(dev, intBuff, size + 1) != size + 1\r\n");
#endif
		return -1;
	}

#ifdef DEBUG
    printf("comm.c: i2cMem8Write success\r\n");
#endif
	return 0;
}



