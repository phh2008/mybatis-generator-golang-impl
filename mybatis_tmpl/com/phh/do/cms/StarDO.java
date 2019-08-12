package com.phh.do.cms;

import java.util.Date;

import com.phh.do.BaseDO;

/**
 * 点赞表
 * 表名：cms_star
 *
 * @author 彭会会
 * @date 2019-08-12
 */
public class StarDO extends BaseDO {
	private static final long serialVersionUID = 1565586068602765100L;


    /**
     * 用户ID
     */
    private Long userId;

    /**
     * 文章ID
     */
    private Long articleId;

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
     * 用户ID
     */
    public void setUserId(Long userId) {
        this.userId = userId;
    }
    /**
     * 用户ID
     */
    public Long getUserId() {
        return userId;
    }

    /**
     * 文章ID
     */
    public void setArticleId(Long articleId) {
        this.articleId = articleId;
    }
    /**
     * 文章ID
     */
    public Long getArticleId() {
        return articleId;
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
