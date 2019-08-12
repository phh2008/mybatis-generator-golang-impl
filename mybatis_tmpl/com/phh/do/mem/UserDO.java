package com.phh.do.mem;

import java.util.Date;

import com.phh.do.BaseDO;

/**
 * 用户表
 * 表名：mem_user
 *
 * @author 彭会会
 * @date 2019-08-12
 */
public class UserDO extends BaseDO {
	private static final long serialVersionUID = 1565586068613733700L;


    /**
     * 用户名
     */
    private String userName;

    /**
     * 密码
     */
    private String password;

    /**
     * 昵称
     */
    private String nickName;

    /**
     * 真实姓名
     */
    private String realName;

    /**
     * 手机号
     */
    private String mobile;

    /**
     * 邮箱
     */
    private String email;

    /**
     * 性别：0-保密，1-男，2-女
     */
    private Integer sex;

    /**
     * 用户类型：0-普通用户，1-管理员
     */
    private Integer type;

    /**
     * 用户状态：0-正常
     */
    private Integer status;

    /**
     * 创建时间
     */
    private Date createdAt;

    /**
     * 更新时间
     */
    private Date updatedAt;

    /**
     * 删除状态：1-正常，2-逻辑删除，3-物理删除
     */
    private Integer state;



    /**
     * 用户名
     */
    public void setUserName(String userName) {
        this.userName = userName;
    }
    /**
     * 用户名
     */
    public String getUserName() {
        return userName;
    }

    /**
     * 密码
     */
    public void setPassword(String password) {
        this.password = password;
    }
    /**
     * 密码
     */
    public String getPassword() {
        return password;
    }

    /**
     * 昵称
     */
    public void setNickName(String nickName) {
        this.nickName = nickName;
    }
    /**
     * 昵称
     */
    public String getNickName() {
        return nickName;
    }

    /**
     * 真实姓名
     */
    public void setRealName(String realName) {
        this.realName = realName;
    }
    /**
     * 真实姓名
     */
    public String getRealName() {
        return realName;
    }

    /**
     * 手机号
     */
    public void setMobile(String mobile) {
        this.mobile = mobile;
    }
    /**
     * 手机号
     */
    public String getMobile() {
        return mobile;
    }

    /**
     * 邮箱
     */
    public void setEmail(String email) {
        this.email = email;
    }
    /**
     * 邮箱
     */
    public String getEmail() {
        return email;
    }

    /**
     * 性别：0-保密，1-男，2-女
     */
    public void setSex(Integer sex) {
        this.sex = sex;
    }
    /**
     * 性别：0-保密，1-男，2-女
     */
    public Integer getSex() {
        return sex;
    }

    /**
     * 用户类型：0-普通用户，1-管理员
     */
    public void setType(Integer type) {
        this.type = type;
    }
    /**
     * 用户类型：0-普通用户，1-管理员
     */
    public Integer getType() {
        return type;
    }

    /**
     * 用户状态：0-正常
     */
    public void setStatus(Integer status) {
        this.status = status;
    }
    /**
     * 用户状态：0-正常
     */
    public Integer getStatus() {
        return status;
    }

    /**
     * 创建时间
     */
    public void setCreatedAt(Date createdAt) {
        this.createdAt = createdAt;
    }
    /**
     * 创建时间
     */
    public Date getCreatedAt() {
        return createdAt;
    }

    /**
     * 更新时间
     */
    public void setUpdatedAt(Date updatedAt) {
        this.updatedAt = updatedAt;
    }
    /**
     * 更新时间
     */
    public Date getUpdatedAt() {
        return updatedAt;
    }

    /**
     * 删除状态：1-正常，2-逻辑删除，3-物理删除
     */
    public void setState(Integer state) {
        this.state = state;
    }
    /**
     * 删除状态：1-正常，2-逻辑删除，3-物理删除
     */
    public Integer getState() {
        return state;
    }


}
